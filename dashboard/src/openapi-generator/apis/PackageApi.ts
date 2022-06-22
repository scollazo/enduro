/* tslint:disable */
/* eslint-disable */
/**
 * Enduro API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import {
    PackageBulkNotAvailableResponseBody,
    PackageBulkNotAvailableResponseBodyFromJSON,
    PackageBulkNotAvailableResponseBodyToJSON,
    PackageBulkNotValidResponseBody,
    PackageBulkNotValidResponseBodyFromJSON,
    PackageBulkNotValidResponseBodyToJSON,
    PackageBulkRequestBody,
    PackageBulkRequestBodyFromJSON,
    PackageBulkRequestBodyToJSON,
    PackageBulkResponseBody,
    PackageBulkResponseBodyFromJSON,
    PackageBulkResponseBodyToJSON,
    PackageBulkStatusResponseBody,
    PackageBulkStatusResponseBodyFromJSON,
    PackageBulkStatusResponseBodyToJSON,
    PackageCancelNotFoundResponseBody,
    PackageCancelNotFoundResponseBodyFromJSON,
    PackageCancelNotFoundResponseBodyToJSON,
    PackageCancelNotRunningResponseBody,
    PackageCancelNotRunningResponseBodyFromJSON,
    PackageCancelNotRunningResponseBodyToJSON,
    PackageConfirmNotAvailableResponseBody,
    PackageConfirmNotAvailableResponseBodyFromJSON,
    PackageConfirmNotAvailableResponseBodyToJSON,
    PackageConfirmNotValidResponseBody,
    PackageConfirmNotValidResponseBodyFromJSON,
    PackageConfirmNotValidResponseBodyToJSON,
    PackageDeleteNotFoundResponseBody,
    PackageDeleteNotFoundResponseBodyFromJSON,
    PackageDeleteNotFoundResponseBodyToJSON,
    PackageListResponseBody,
    PackageListResponseBodyFromJSON,
    PackageListResponseBodyToJSON,
    PackageMonitorResponseBody,
    PackageMonitorResponseBodyFromJSON,
    PackageMonitorResponseBodyToJSON,
    PackagePreservationActionsNotFoundResponseBody,
    PackagePreservationActionsNotFoundResponseBodyFromJSON,
    PackagePreservationActionsNotFoundResponseBodyToJSON,
    PackagePreservationActionsResponseBody,
    PackagePreservationActionsResponseBodyFromJSON,
    PackagePreservationActionsResponseBodyToJSON,
    PackageRejectNotAvailableResponseBody,
    PackageRejectNotAvailableResponseBodyFromJSON,
    PackageRejectNotAvailableResponseBodyToJSON,
    PackageRejectNotValidResponseBody,
    PackageRejectNotValidResponseBodyFromJSON,
    PackageRejectNotValidResponseBodyToJSON,
    PackageRetryNotFoundResponseBody,
    PackageRetryNotFoundResponseBodyFromJSON,
    PackageRetryNotFoundResponseBodyToJSON,
    PackageRetryNotRunningResponseBody,
    PackageRetryNotRunningResponseBodyFromJSON,
    PackageRetryNotRunningResponseBodyToJSON,
    PackageShowNotFoundResponseBody,
    PackageShowNotFoundResponseBodyFromJSON,
    PackageShowNotFoundResponseBodyToJSON,
    PackageShowResponseBody,
    PackageShowResponseBodyFromJSON,
    PackageShowResponseBodyToJSON,
    PackageWorkflowNotFoundResponseBody,
    PackageWorkflowNotFoundResponseBodyFromJSON,
    PackageWorkflowNotFoundResponseBodyToJSON,
    PackageWorkflowResponseBody,
    PackageWorkflowResponseBodyFromJSON,
    PackageWorkflowResponseBodyToJSON,
} from '../models';

export interface PackageBulkRequest {
    bulkRequestBody: PackageBulkRequestBody;
}

export interface PackageCancelRequest {
    id: number;
}

export interface PackageConfirmRequest {
    id: number;
}

export interface PackageDeleteRequest {
    id: number;
}

export interface PackageListRequest {
    name?: string;
    aipId?: string;
    earliestCreatedTime?: Date;
    latestCreatedTime?: Date;
    status?: PackageListStatusEnum;
    cursor?: string;
}

export interface PackagePreservationActionsRequest {
    id: number;
}

export interface PackageRejectRequest {
    id: number;
}

export interface PackageRetryRequest {
    id: number;
}

export interface PackageShowRequest {
    id: number;
}

export interface PackageWorkflowRequest {
    id: number;
}

/**
 * PackageApi - interface
 * 
 * @export
 * @interface PackageApiInterface
 */
export interface PackageApiInterface {
    /**
     * Bulk operations (retry, cancel...).
     * @summary bulk package
     * @param {PackageBulkRequestBody} bulkRequestBody 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PackageApiInterface
     */
    packageBulkRaw(requestParameters: PackageBulkRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<PackageBulkResponseBody>>;

    /**
     * Bulk operations (retry, cancel...).
     * bulk package
     */
    packageBulk(requestParameters: PackageBulkRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<PackageBulkResponseBody>;

    /**
     * Retrieve status of current bulk operation.
     * @summary bulk_status package
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PackageApiInterface
     */
    packageBulkStatusRaw(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<PackageBulkStatusResponseBody>>;

    /**
     * Retrieve status of current bulk operation.
     * bulk_status package
     */
    packageBulkStatus(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<PackageBulkStatusResponseBody>;

    /**
     * Cancel package processing by ID
     * @summary cancel package
     * @param {number} id Identifier of package to remove
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PackageApiInterface
     */
    packageCancelRaw(requestParameters: PackageCancelRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>>;

    /**
     * Cancel package processing by ID
     * cancel package
     */
    packageCancel(requestParameters: PackageCancelRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void>;

    /**
     * Signal the package has been reviewed and accepted
     * @summary confirm package
     * @param {number} id Identifier of package to look up
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PackageApiInterface
     */
    packageConfirmRaw(requestParameters: PackageConfirmRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>>;

    /**
     * Signal the package has been reviewed and accepted
     * confirm package
     */
    packageConfirm(requestParameters: PackageConfirmRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void>;

    /**
     * Delete package by ID
     * @summary delete package
     * @param {number} id Identifier of package to delete
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PackageApiInterface
     */
    packageDeleteRaw(requestParameters: PackageDeleteRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>>;

    /**
     * Delete package by ID
     * delete package
     */
    packageDelete(requestParameters: PackageDeleteRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void>;

    /**
     * List all stored packages
     * @summary list package
     * @param {string} [name] 
     * @param {string} [aipId] 
     * @param {Date} [earliestCreatedTime] 
     * @param {Date} [latestCreatedTime] 
     * @param {'new' | 'in progress' | 'done' | 'error' | 'unknown' | 'queued' | 'pending' | 'abandoned'} [status] 
     * @param {string} [cursor] Pagination cursor
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PackageApiInterface
     */
    packageListRaw(requestParameters: PackageListRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<PackageListResponseBody>>;

    /**
     * List all stored packages
     * list package
     */
    packageList(requestParameters: PackageListRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<PackageListResponseBody>;

    /**
     * 
     * @summary monitor package
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PackageApiInterface
     */
    packageMonitorRaw(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>>;

    /**
     * monitor package
     */
    packageMonitor(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void>;

    /**
     * List all preservation actions by ID
     * @summary preservation-actions package
     * @param {number} id Identifier of package to look up
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PackageApiInterface
     */
    packagePreservationActionsRaw(requestParameters: PackagePreservationActionsRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<PackagePreservationActionsResponseBody>>;

    /**
     * List all preservation actions by ID
     * preservation-actions package
     */
    packagePreservationActions(requestParameters: PackagePreservationActionsRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<PackagePreservationActionsResponseBody>;

    /**
     * Signal the package has been reviewed and rejected
     * @summary reject package
     * @param {number} id Identifier of package to look up
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PackageApiInterface
     */
    packageRejectRaw(requestParameters: PackageRejectRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>>;

    /**
     * Signal the package has been reviewed and rejected
     * reject package
     */
    packageReject(requestParameters: PackageRejectRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void>;

    /**
     * Retry package processing by ID
     * @summary retry package
     * @param {number} id Identifier of package to retry
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PackageApiInterface
     */
    packageRetryRaw(requestParameters: PackageRetryRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>>;

    /**
     * Retry package processing by ID
     * retry package
     */
    packageRetry(requestParameters: PackageRetryRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void>;

    /**
     * Show package by ID
     * @summary show package
     * @param {number} id Identifier of package to show
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PackageApiInterface
     */
    packageShowRaw(requestParameters: PackageShowRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<PackageShowResponseBody>>;

    /**
     * Show package by ID
     * show package
     */
    packageShow(requestParameters: PackageShowRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<PackageShowResponseBody>;

    /**
     * Retrieve workflow status by ID
     * @summary workflow package
     * @param {number} id Identifier of package to look up
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PackageApiInterface
     */
    packageWorkflowRaw(requestParameters: PackageWorkflowRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<PackageWorkflowResponseBody>>;

    /**
     * Retrieve workflow status by ID
     * workflow package
     */
    packageWorkflow(requestParameters: PackageWorkflowRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<PackageWorkflowResponseBody>;

}

/**
 * 
 */
export class PackageApi extends runtime.BaseAPI implements PackageApiInterface {

    /**
     * Bulk operations (retry, cancel...).
     * bulk package
     */
    async packageBulkRaw(requestParameters: PackageBulkRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<PackageBulkResponseBody>> {
        if (requestParameters.bulkRequestBody === null || requestParameters.bulkRequestBody === undefined) {
            throw new runtime.RequiredError('bulkRequestBody','Required parameter requestParameters.bulkRequestBody was null or undefined when calling packageBulk.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/package/bulk`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: PackageBulkRequestBodyToJSON(requestParameters.bulkRequestBody),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PackageBulkResponseBodyFromJSON(jsonValue));
    }

    /**
     * Bulk operations (retry, cancel...).
     * bulk package
     */
    async packageBulk(requestParameters: PackageBulkRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<PackageBulkResponseBody> {
        const response = await this.packageBulkRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Retrieve status of current bulk operation.
     * bulk_status package
     */
    async packageBulkStatusRaw(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<PackageBulkStatusResponseBody>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/package/bulk`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PackageBulkStatusResponseBodyFromJSON(jsonValue));
    }

    /**
     * Retrieve status of current bulk operation.
     * bulk_status package
     */
    async packageBulkStatus(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<PackageBulkStatusResponseBody> {
        const response = await this.packageBulkStatusRaw(initOverrides);
        return await response.value();
    }

    /**
     * Cancel package processing by ID
     * cancel package
     */
    async packageCancelRaw(requestParameters: PackageCancelRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling packageCancel.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/package/{id}/cancel`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Cancel package processing by ID
     * cancel package
     */
    async packageCancel(requestParameters: PackageCancelRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void> {
        await this.packageCancelRaw(requestParameters, initOverrides);
    }

    /**
     * Signal the package has been reviewed and accepted
     * confirm package
     */
    async packageConfirmRaw(requestParameters: PackageConfirmRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling packageConfirm.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/package/{id}/confirm`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Signal the package has been reviewed and accepted
     * confirm package
     */
    async packageConfirm(requestParameters: PackageConfirmRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void> {
        await this.packageConfirmRaw(requestParameters, initOverrides);
    }

    /**
     * Delete package by ID
     * delete package
     */
    async packageDeleteRaw(requestParameters: PackageDeleteRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling packageDelete.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/package/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Delete package by ID
     * delete package
     */
    async packageDelete(requestParameters: PackageDeleteRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void> {
        await this.packageDeleteRaw(requestParameters, initOverrides);
    }

    /**
     * List all stored packages
     * list package
     */
    async packageListRaw(requestParameters: PackageListRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<PackageListResponseBody>> {
        const queryParameters: any = {};

        if (requestParameters.name !== undefined) {
            queryParameters['name'] = requestParameters.name;
        }

        if (requestParameters.aipId !== undefined) {
            queryParameters['aip_id'] = requestParameters.aipId;
        }

        if (requestParameters.earliestCreatedTime !== undefined) {
            queryParameters['earliest_created_time'] = (requestParameters.earliestCreatedTime as any).toISOString();
        }

        if (requestParameters.latestCreatedTime !== undefined) {
            queryParameters['latest_created_time'] = (requestParameters.latestCreatedTime as any).toISOString();
        }

        if (requestParameters.status !== undefined) {
            queryParameters['status'] = requestParameters.status;
        }

        if (requestParameters.cursor !== undefined) {
            queryParameters['cursor'] = requestParameters.cursor;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/package`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PackageListResponseBodyFromJSON(jsonValue));
    }

    /**
     * List all stored packages
     * list package
     */
    async packageList(requestParameters: PackageListRequest = {}, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<PackageListResponseBody> {
        const response = await this.packageListRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * monitor package
     */
    async packageMonitorRaw(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/package/monitor`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * monitor package
     */
    async packageMonitor(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void> {
        await this.packageMonitorRaw(initOverrides);
    }

    /**
     * List all preservation actions by ID
     * preservation-actions package
     */
    async packagePreservationActionsRaw(requestParameters: PackagePreservationActionsRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<PackagePreservationActionsResponseBody>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling packagePreservationActions.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/package/{id}/preservation-actions`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PackagePreservationActionsResponseBodyFromJSON(jsonValue));
    }

    /**
     * List all preservation actions by ID
     * preservation-actions package
     */
    async packagePreservationActions(requestParameters: PackagePreservationActionsRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<PackagePreservationActionsResponseBody> {
        const response = await this.packagePreservationActionsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Signal the package has been reviewed and rejected
     * reject package
     */
    async packageRejectRaw(requestParameters: PackageRejectRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling packageReject.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/package/{id}/reject`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Signal the package has been reviewed and rejected
     * reject package
     */
    async packageReject(requestParameters: PackageRejectRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void> {
        await this.packageRejectRaw(requestParameters, initOverrides);
    }

    /**
     * Retry package processing by ID
     * retry package
     */
    async packageRetryRaw(requestParameters: PackageRetryRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling packageRetry.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/package/{id}/retry`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Retry package processing by ID
     * retry package
     */
    async packageRetry(requestParameters: PackageRetryRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void> {
        await this.packageRetryRaw(requestParameters, initOverrides);
    }

    /**
     * Show package by ID
     * show package
     */
    async packageShowRaw(requestParameters: PackageShowRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<PackageShowResponseBody>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling packageShow.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/package/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PackageShowResponseBodyFromJSON(jsonValue));
    }

    /**
     * Show package by ID
     * show package
     */
    async packageShow(requestParameters: PackageShowRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<PackageShowResponseBody> {
        const response = await this.packageShowRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Retrieve workflow status by ID
     * workflow package
     */
    async packageWorkflowRaw(requestParameters: PackageWorkflowRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<PackageWorkflowResponseBody>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling packageWorkflow.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/package/{id}/workflow`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PackageWorkflowResponseBodyFromJSON(jsonValue));
    }

    /**
     * Retrieve workflow status by ID
     * workflow package
     */
    async packageWorkflow(requestParameters: PackageWorkflowRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<PackageWorkflowResponseBody> {
        const response = await this.packageWorkflowRaw(requestParameters, initOverrides);
        return await response.value();
    }

}

/**
 * @export
 */
export const PackageListStatusEnum = {
    New: 'new',
    InProgress: 'in progress',
    Done: 'done',
    Error: 'error',
    Unknown: 'unknown',
    Queued: 'queued',
    Pending: 'pending',
    Abandoned: 'abandoned'
} as const;
export type PackageListStatusEnum = typeof PackageListStatusEnum[keyof typeof PackageListStatusEnum];
