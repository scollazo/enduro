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
 * Collection not found
 * @export
 * @interface PipelineShowNotFoundResponseBody
 */
export interface PipelineShowNotFoundResponseBody {
    /**
     * Identifier of missing collection
     * @type {number}
     * @memberof PipelineShowNotFoundResponseBody
     */
    id: number;
    /**
     * Message of error
     * @type {string}
     * @memberof PipelineShowNotFoundResponseBody
     */
    message: string;
}

export function PipelineShowNotFoundResponseBodyFromJSON(json: any): PipelineShowNotFoundResponseBody {
    return PipelineShowNotFoundResponseBodyFromJSONTyped(json, false);
}

export function PipelineShowNotFoundResponseBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): PipelineShowNotFoundResponseBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'message': json['message'],
    };
}

export function PipelineShowNotFoundResponseBodyToJSON(value?: PipelineShowNotFoundResponseBody | null): any {
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


