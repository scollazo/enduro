/* tslint:disable */
/* eslint-disable */
/**
 * Enduro API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 0.0.1
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  ConfirmRequestBody,
  EnduroPackagePreservationActions,
  EnduroStoredPackage,
  ListResponseBody,
  MonitorEvent,
  MoveStatusResult,
  PackageNotFound,
} from '../models/index';
import {
    ConfirmRequestBodyFromJSON,
    ConfirmRequestBodyToJSON,
    EnduroPackagePreservationActionsFromJSON,
    EnduroPackagePreservationActionsToJSON,
    EnduroStoredPackageFromJSON,
    EnduroStoredPackageToJSON,
    ListResponseBodyFromJSON,
    ListResponseBodyToJSON,
    MonitorEventFromJSON,
    MonitorEventToJSON,
    MoveStatusResultFromJSON,
    MoveStatusResultToJSON,
    PackageNotFoundFromJSON,
    PackageNotFoundToJSON,
} from '../models/index';

export interface PackageConfirmRequest {
    id: number;
    confirmRequestBody: ConfirmRequestBody;
}

export interface PackageListRequest {
    name?: string;
    aipId?: string;
    earliestCreatedTime?: Date;
    latestCreatedTime?: Date;
    locationId?: string;
    status?: PackageListStatusEnum;
    cursor?: string;
}

export interface PackageMonitorRequest {
    enduroWsTicket?: string;
}

export interface PackageMoveRequest {
    id: number;
    confirmRequestBody: ConfirmRequestBody;
}

export interface PackageMoveStatusRequest {
    id: number;
}

export interface PackagePreservationActionsRequest {
    id: number;
}

export interface PackageRejectRequest {
    id: number;
}

export interface PackageShowRequest {
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
     * Signal the package has been reviewed and accepted
     * @summary confirm package
     * @param {number} id Identifier of package to look up
     * @param {ConfirmRequestBody} confirmRequestBody 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PackageApiInterface
     */
    packageConfirmRaw(requestParameters: PackageConfirmRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>>;

    /**
     * Signal the package has been reviewed and accepted
     * confirm package
     */
    packageConfirm(requestParameters: PackageConfirmRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void>;

    /**
     * List all stored packages
     * @summary list package
     * @param {string} [name] 
     * @param {string} [aipId] Identifier of AIP
     * @param {Date} [earliestCreatedTime] 
     * @param {Date} [latestCreatedTime] 
     * @param {string} [locationId] Identifier of storage location
     * @param {'new' | 'in progress' | 'done' | 'error' | 'unknown' | 'queued' | 'pending' | 'abandoned'} [status] 
     * @param {string} [cursor] Pagination cursor
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PackageApiInterface
     */
    packageListRaw(requestParameters: PackageListRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ListResponseBody>>;

    /**
     * List all stored packages
     * list package
     */
    packageList(requestParameters: PackageListRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ListResponseBody>;

    /**
     * 
     * @summary monitor package
     * @param {string} [enduroWsTicket] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PackageApiInterface
     */
    packageMonitorRaw(requestParameters: PackageMonitorRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>>;

    /**
     * monitor package
     */
    packageMonitor(requestParameters: PackageMonitorRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void>;

    /**
     * Request access to the /monitor WebSocket.
     * @summary monitor_request package
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PackageApiInterface
     */
    packageMonitorRequestRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>>;

    /**
     * Request access to the /monitor WebSocket.
     * monitor_request package
     */
    packageMonitorRequest(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void>;

    /**
     * Move a package to a permanent storage location
     * @summary move package
     * @param {number} id Identifier of package to move
     * @param {ConfirmRequestBody} confirmRequestBody 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PackageApiInterface
     */
    packageMoveRaw(requestParameters: PackageMoveRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>>;

    /**
     * Move a package to a permanent storage location
     * move package
     */
    packageMove(requestParameters: PackageMoveRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void>;

    /**
     * Retrieve the status of a permanent storage location move of the package
     * @summary move_status package
     * @param {number} id Identifier of package to move
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PackageApiInterface
     */
    packageMoveStatusRaw(requestParameters: PackageMoveStatusRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<MoveStatusResult>>;

    /**
     * Retrieve the status of a permanent storage location move of the package
     * move_status package
     */
    packageMoveStatus(requestParameters: PackageMoveStatusRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<MoveStatusResult>;

    /**
     * List all preservation actions by ID
     * @summary preservation_actions package
     * @param {number} id Identifier of package to look up
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PackageApiInterface
     */
    packagePreservationActionsRaw(requestParameters: PackagePreservationActionsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<EnduroPackagePreservationActions>>;

    /**
     * List all preservation actions by ID
     * preservation_actions package
     */
    packagePreservationActions(requestParameters: PackagePreservationActionsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<EnduroPackagePreservationActions>;

    /**
     * Signal the package has been reviewed and rejected
     * @summary reject package
     * @param {number} id Identifier of package to look up
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PackageApiInterface
     */
    packageRejectRaw(requestParameters: PackageRejectRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>>;

    /**
     * Signal the package has been reviewed and rejected
     * reject package
     */
    packageReject(requestParameters: PackageRejectRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void>;

    /**
     * Show package by ID
     * @summary show package
     * @param {number} id Identifier of package to show
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PackageApiInterface
     */
    packageShowRaw(requestParameters: PackageShowRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<EnduroStoredPackage>>;

    /**
     * Show package by ID
     * show package
     */
    packageShow(requestParameters: PackageShowRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<EnduroStoredPackage>;

}

/**
 * 
 */
export class PackageApi extends runtime.BaseAPI implements PackageApiInterface {

    /**
     * Signal the package has been reviewed and accepted
     * confirm package
     */
    async packageConfirmRaw(requestParameters: PackageConfirmRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling packageConfirm.');
        }

        if (requestParameters.confirmRequestBody === null || requestParameters.confirmRequestBody === undefined) {
            throw new runtime.RequiredError('confirmRequestBody','Required parameter requestParameters.confirmRequestBody was null or undefined when calling packageConfirm.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.accessToken) {
            // oauth required
            headerParameters["Authorization"] = await this.configuration.accessToken("oauth2_header_Authorization", []);
        }

        const response = await this.request({
            path: `/package/{id}/confirm`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ConfirmRequestBodyToJSON(requestParameters.confirmRequestBody),
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Signal the package has been reviewed and accepted
     * confirm package
     */
    async packageConfirm(requestParameters: PackageConfirmRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.packageConfirmRaw(requestParameters, initOverrides);
    }

    /**
     * List all stored packages
     * list package
     */
    async packageListRaw(requestParameters: PackageListRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ListResponseBody>> {
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

        if (requestParameters.locationId !== undefined) {
            queryParameters['location_id'] = requestParameters.locationId;
        }

        if (requestParameters.status !== undefined) {
            queryParameters['status'] = requestParameters.status;
        }

        if (requestParameters.cursor !== undefined) {
            queryParameters['cursor'] = requestParameters.cursor;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            // oauth required
            headerParameters["Authorization"] = await this.configuration.accessToken("oauth2_header_Authorization", []);
        }

        const response = await this.request({
            path: `/package`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ListResponseBodyFromJSON(jsonValue));
    }

    /**
     * List all stored packages
     * list package
     */
    async packageList(requestParameters: PackageListRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ListResponseBody> {
        const response = await this.packageListRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * monitor package
     */
    async packageMonitorRaw(requestParameters: PackageMonitorRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
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
    async packageMonitor(requestParameters: PackageMonitorRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.packageMonitorRaw(requestParameters, initOverrides);
    }

    /**
     * Request access to the /monitor WebSocket.
     * monitor_request package
     */
    async packageMonitorRequestRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            // oauth required
            headerParameters["Authorization"] = await this.configuration.accessToken("oauth2_header_Authorization", []);
        }

        const response = await this.request({
            path: `/package/monitor`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Request access to the /monitor WebSocket.
     * monitor_request package
     */
    async packageMonitorRequest(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.packageMonitorRequestRaw(initOverrides);
    }

    /**
     * Move a package to a permanent storage location
     * move package
     */
    async packageMoveRaw(requestParameters: PackageMoveRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling packageMove.');
        }

        if (requestParameters.confirmRequestBody === null || requestParameters.confirmRequestBody === undefined) {
            throw new runtime.RequiredError('confirmRequestBody','Required parameter requestParameters.confirmRequestBody was null or undefined when calling packageMove.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.accessToken) {
            // oauth required
            headerParameters["Authorization"] = await this.configuration.accessToken("oauth2_header_Authorization", []);
        }

        const response = await this.request({
            path: `/package/{id}/move`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ConfirmRequestBodyToJSON(requestParameters.confirmRequestBody),
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Move a package to a permanent storage location
     * move package
     */
    async packageMove(requestParameters: PackageMoveRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.packageMoveRaw(requestParameters, initOverrides);
    }

    /**
     * Retrieve the status of a permanent storage location move of the package
     * move_status package
     */
    async packageMoveStatusRaw(requestParameters: PackageMoveStatusRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<MoveStatusResult>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling packageMoveStatus.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            // oauth required
            headerParameters["Authorization"] = await this.configuration.accessToken("oauth2_header_Authorization", []);
        }

        const response = await this.request({
            path: `/package/{id}/move`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => MoveStatusResultFromJSON(jsonValue));
    }

    /**
     * Retrieve the status of a permanent storage location move of the package
     * move_status package
     */
    async packageMoveStatus(requestParameters: PackageMoveStatusRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<MoveStatusResult> {
        const response = await this.packageMoveStatusRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List all preservation actions by ID
     * preservation_actions package
     */
    async packagePreservationActionsRaw(requestParameters: PackagePreservationActionsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<EnduroPackagePreservationActions>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling packagePreservationActions.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            // oauth required
            headerParameters["Authorization"] = await this.configuration.accessToken("oauth2_header_Authorization", []);
        }

        const response = await this.request({
            path: `/package/{id}/preservation-actions`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => EnduroPackagePreservationActionsFromJSON(jsonValue));
    }

    /**
     * List all preservation actions by ID
     * preservation_actions package
     */
    async packagePreservationActions(requestParameters: PackagePreservationActionsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<EnduroPackagePreservationActions> {
        const response = await this.packagePreservationActionsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Signal the package has been reviewed and rejected
     * reject package
     */
    async packageRejectRaw(requestParameters: PackageRejectRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling packageReject.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            // oauth required
            headerParameters["Authorization"] = await this.configuration.accessToken("oauth2_header_Authorization", []);
        }

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
    async packageReject(requestParameters: PackageRejectRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.packageRejectRaw(requestParameters, initOverrides);
    }

    /**
     * Show package by ID
     * show package
     */
    async packageShowRaw(requestParameters: PackageShowRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<EnduroStoredPackage>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling packageShow.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            // oauth required
            headerParameters["Authorization"] = await this.configuration.accessToken("oauth2_header_Authorization", []);
        }

        const response = await this.request({
            path: `/package/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => EnduroStoredPackageFromJSON(jsonValue));
    }

    /**
     * Show package by ID
     * show package
     */
    async packageShow(requestParameters: PackageShowRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<EnduroStoredPackage> {
        const response = await this.packageShowRaw(requestParameters, initOverrides);
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
