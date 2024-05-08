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
 * @interface SubmitRequestBody
 */
export interface SubmitRequestBody {
    /**
     * 
     * @type {string}
     * @memberof SubmitRequestBody
     */
    name: string;
}

/**
 * Check if a given object implements the SubmitRequestBody interface.
 */
export function instanceOfSubmitRequestBody(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "name" in value;

    return isInstance;
}

export function SubmitRequestBodyFromJSON(json: any): SubmitRequestBody {
    return SubmitRequestBodyFromJSONTyped(json, false);
}

export function SubmitRequestBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): SubmitRequestBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': json['name'],
    };
}

export function SubmitRequestBodyToJSON(value?: SubmitRequestBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
    };
}

