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

import { exists, mapValues } from '../runtime';
/**
 * Package not found
 * @export
 * @interface PackageWorkflowNotFoundResponseBody
 */
export interface PackageWorkflowNotFoundResponseBody {
    /**
     * Identifier of missing package
     * @type {number}
     * @memberof PackageWorkflowNotFoundResponseBody
     */
    id: number;
    /**
     * Message of error
     * @type {string}
     * @memberof PackageWorkflowNotFoundResponseBody
     */
    message: string;
}

export function PackageWorkflowNotFoundResponseBodyFromJSON(json: any): PackageWorkflowNotFoundResponseBody {
    return PackageWorkflowNotFoundResponseBodyFromJSONTyped(json, false);
}

export function PackageWorkflowNotFoundResponseBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): PackageWorkflowNotFoundResponseBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'message': json['message'],
    };
}

export function PackageWorkflowNotFoundResponseBodyToJSON(value?: PackageWorkflowNotFoundResponseBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'message': value.message,
    };
}

