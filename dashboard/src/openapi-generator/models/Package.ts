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
 * @interface Package
 */
export interface Package {
    /**
     * 
     * @type {string}
     * @memberof Package
     */
    aipId: string;
    /**
     * Creation datetime
     * @type {Date}
     * @memberof Package
     */
    createdAt: Date;
    /**
     * Identifier of storage location
     * @type {string}
     * @memberof Package
     */
    locationId?: string;
    /**
     * 
     * @type {string}
     * @memberof Package
     */
    name: string;
    /**
     * 
     * @type {string}
     * @memberof Package
     */
    objectKey: string;
    /**
     * Status of the package
     * @type {string}
     * @memberof Package
     */
    status: PackageStatusEnum;
}


/**
 * @export
 */
export const PackageStatusEnum = {
    Unspecified: 'unspecified',
    InReview: 'in_review',
    Rejected: 'rejected',
    Stored: 'stored',
    Moving: 'moving'
} as const;
export type PackageStatusEnum = typeof PackageStatusEnum[keyof typeof PackageStatusEnum];


/**
 * Check if a given object implements the Package interface.
 */
export function instanceOfPackage(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "aipId" in value;
    isInstance = isInstance && "createdAt" in value;
    isInstance = isInstance && "name" in value;
    isInstance = isInstance && "objectKey" in value;
    isInstance = isInstance && "status" in value;

    return isInstance;
}

export function PackageFromJSON(json: any): Package {
    return PackageFromJSONTyped(json, false);
}

export function PackageFromJSONTyped(json: any, ignoreDiscriminator: boolean): Package {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'aipId': json['aip_id'],
        'createdAt': (new Date(json['created_at'])),
        'locationId': !exists(json, 'location_id') ? undefined : json['location_id'],
        'name': json['name'],
        'objectKey': json['object_key'],
        'status': json['status'],
    };
}

export function PackageToJSON(value?: Package | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'aip_id': value.aipId,
        'created_at': (value.createdAt.toISOString()),
        'location_id': value.locationId,
        'name': value.name,
        'object_key': value.objectKey,
        'status': value.status,
    };
}

