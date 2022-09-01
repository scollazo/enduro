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
 * Error returned when the Content-Type header does not define a multipart request. (default view)
 * @export
 * @interface UploadUploadInvalidMediaTypeResponseBody
 */
export interface UploadUploadInvalidMediaTypeResponseBody {
    /**
     * Is the error a server-side fault?
     * @type {boolean}
     * @memberof UploadUploadInvalidMediaTypeResponseBody
     */
    fault: boolean;
    /**
     * ID is a unique identifier for this particular occurrence of the problem.
     * @type {string}
     * @memberof UploadUploadInvalidMediaTypeResponseBody
     */
    id: string;
    /**
     * Message is a human-readable explanation specific to this occurrence of the problem.
     * @type {string}
     * @memberof UploadUploadInvalidMediaTypeResponseBody
     */
    message: string;
    /**
     * Name is the name of this class of errors.
     * @type {string}
     * @memberof UploadUploadInvalidMediaTypeResponseBody
     */
    name: string;
    /**
     * Is the error temporary?
     * @type {boolean}
     * @memberof UploadUploadInvalidMediaTypeResponseBody
     */
    temporary: boolean;
    /**
     * Is the error a timeout?
     * @type {boolean}
     * @memberof UploadUploadInvalidMediaTypeResponseBody
     */
    timeout: boolean;
}

export function UploadUploadInvalidMediaTypeResponseBodyFromJSON(json: any): UploadUploadInvalidMediaTypeResponseBody {
    return UploadUploadInvalidMediaTypeResponseBodyFromJSONTyped(json, false);
}

export function UploadUploadInvalidMediaTypeResponseBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): UploadUploadInvalidMediaTypeResponseBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'fault': json['fault'],
        'id': json['id'],
        'message': json['message'],
        'name': json['name'],
        'temporary': json['temporary'],
        'timeout': json['timeout'],
    };
}

export function UploadUploadInvalidMediaTypeResponseBodyToJSON(value?: UploadUploadInvalidMediaTypeResponseBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'fault': value.fault,
        'id': value.id,
        'message': value.message,
        'name': value.name,
        'temporary': value.temporary,
        'timeout': value.timeout,
    };
}

