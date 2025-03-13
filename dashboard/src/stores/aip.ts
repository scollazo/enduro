import { acceptHMRUpdate, defineStore } from "pinia";

import { api, client } from "@/client";
import { ResponseError, StorageListAipsStatusEnum } from "@/openapi-generator";
import router from "@/router";
import { useLayoutStore } from "@/stores/layout";

const defaultPageSize = 20;

export interface Pager {
  // maxPages is the maximum number of page links to show in the pager.
  readonly maxPages: number;

  current: number;
  first: number;
  last: number;
  total: number;
  pages: Array<number>;
}

export const useAipStore = defineStore("aip", {
  state: () => ({
    // AIP currently displayed.
    current: null as api.EnduroStorageAip | null,

    // A list of AIPs shown during searches.
    aips: [] as Array<api.EnduroStorageAip>,

    // Page is a subset of the total AIP list.
    page: { limit: defaultPageSize } as api.EnduroPage,

    // Pager contains a list of page numbers to show in the pager.
    pager: { maxPages: 7 } as Pager,

    filters: {
      name: "" as string | undefined,
      status: "" as StorageListAipsStatusEnum | undefined,
      earliestCreatedTime: undefined as Date | undefined,
      latestCreatedTime: undefined as Date | undefined,
    },
  }),
  getters: {
    hasNextPage(): boolean {
      return this.page.offset + this.page.limit < this.page.total;
    },
    hasPrevPage(): boolean {
      return this.page.offset > 0;
    },
    lastResultOnPage(): number {
      let i = this.page.offset + this.page.limit;
      if (i > this.page.total) {
        i = this.page.total;
      }
      return i;
    },
  },
  actions: {
    async fetchCurrent(id: string) {
      this.current = await client.storage.storageShowAip({ uuid: id });

      // Update breadcrumb. TODO: should this be done in the component?
      const layoutStore = useLayoutStore();
      layoutStore.updateBreadcrumb([
        { text: "Storage" },
        { route: router.resolve("/storage/aips/"), text: "AIPs" },
        { text: this.current.name },
      ]);
    },
    async fetchAips(page: number) {
      return client.storage
        .storageListAips({
          offset: page > 1 ? (page - 1) * this.page.limit : undefined,
          limit: this.page?.limit || undefined,
          name: this.filters.name,
          status: this.filters.status,
          earliestCreatedTime: this.filters.earliestCreatedTime,
          latestCreatedTime: this.filters.latestCreatedTime,
        })
        .then((resp) => {
          this.aips = resp.items;
          this.page = resp.page;
          this.updatePager();
        })
        .catch(async (err) => {
          this.aips = [];
          this.page = { limit: defaultPageSize, offset: 0, total: 0 };
          this.updatePager();

          if (err instanceof ResponseError) {
            // An invalid status or time range returns a ResponseError with the
            // error message in the response body (JSON).
            return err.response.text().then((body) => {
              const modelErr = api.ModelErrorFromJSON(JSON.parse(body));
              console.error(
                "API response",
                err.response.status,
                modelErr.message,
              );
              throw new Error(modelErr.message);
            });
          } else if (err instanceof RangeError) {
            // An invalid date parameter (e.g. earliestCreatedTime) returns a
            // RangeError with a message like "invalid date".
            console.error("Range error", err.message);
            throw new Error(err.message);
          } else {
            console.error("Unknown error", err.message);
            throw new Error(err.message);
          }
        });
    },
    nextPage() {
      if (this.hasNextPage) {
        this.fetchAips(this.pager.current + 1);
      }
    },
    prevPage() {
      if (this.hasPrevPage) {
        this.fetchAips(this.pager.current - 1);
      }
    },
    updatePager(): void {
      const pgr = this.pager;
      pgr.total = Math.ceil(this.page.total / this.page.limit);
      pgr.current = Math.floor(this.page.offset / this.page.limit) + 1;

      let first = 1;
      const count = pgr.total < pgr.maxPages ? pgr.total : pgr.maxPages;
      const half = Math.floor(pgr.maxPages / 2);
      if (pgr.current > half + 1) {
        if (pgr.total - pgr.current < half) {
          first = pgr.total - count + 1;
        } else {
          first = pgr.current - half;
        }
      }
      pgr.first = first;
      pgr.last = first + count - 1;

      pgr.pages = new Array(count);
      for (let i = 0; i < count; i++) {
        pgr.pages[i] = i + first;
      }
    },
  },
});

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useAipStore, import.meta.hot));
}
