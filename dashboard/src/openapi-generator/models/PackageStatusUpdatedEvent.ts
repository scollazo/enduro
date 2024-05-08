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

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface PackageStatusUpdatedEvent
 */
export interface PackageStatusUpdatedEvent {
    /**
     * Identifier of package
     * @type {number}
     * @memberof PackageStatusUpdatedEvent
     */
    id: number;
    /**
     * 
     * @type {string}
     * @memberof PackageStatusUpdatedEvent
     */
    status: PackageStatusUpdatedEventStatusEnum;
}


/**
 * @export
 */
export const PackageStatusUpdatedEventStatusEnum = {
    New: 'new',
    InProgress: 'in progress',
    Done: 'done',
    Error: 'error',
    Unknown: 'unknown',
    Queued: 'queued',
    Pending: 'pending',
    Abandoned: 'abandoned'
} as const;
export type PackageStatusUpdatedEventStatusEnum = typeof PackageStatusUpdatedEventStatusEnum[keyof typeof PackageStatusUpdatedEventStatusEnum];


/**
 * Check if a given object implements the PackageStatusUpdatedEvent interface.
 */
export function instanceOfPackageStatusUpdatedEvent(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "id" in value;
    isInstance = isInstance && "status" in value;

    return isInstance;
}

export function PackageStatusUpdatedEventFromJSON(json: any): PackageStatusUpdatedEvent {
    return PackageStatusUpdatedEventFromJSONTyped(json, false);
}

export function PackageStatusUpdatedEventFromJSONTyped(json: any, ignoreDiscriminator: boolean): PackageStatusUpdatedEvent {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'status': json['status'],
    };
}

export function PackageStatusUpdatedEventToJSON(value?: PackageStatusUpdatedEvent | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'status': value.status,
    };
}

