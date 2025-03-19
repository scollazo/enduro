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
import type { EnduroIngestSipTask } from './EnduroIngestSipTask';
import {
    EnduroIngestSipTaskFromJSON,
    EnduroIngestSipTaskFromJSONTyped,
    EnduroIngestSipTaskToJSON,
} from './EnduroIngestSipTask';

/**
 * 
 * @export
 * @interface SIPTaskUpdatedEvent
 */
export interface SIPTaskUpdatedEvent {
    /**
     * Identifier of task
     * @type {number}
     * @memberof SIPTaskUpdatedEvent
     */
    id: number;
    /**
     * 
     * @type {EnduroIngestSipTask}
     * @memberof SIPTaskUpdatedEvent
     */
    item: EnduroIngestSipTask;
}

/**
 * Check if a given object implements the SIPTaskUpdatedEvent interface.
 */
export function instanceOfSIPTaskUpdatedEvent(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "id" in value;
    isInstance = isInstance && "item" in value;

    return isInstance;
}

export function SIPTaskUpdatedEventFromJSON(json: any): SIPTaskUpdatedEvent {
    return SIPTaskUpdatedEventFromJSONTyped(json, false);
}

export function SIPTaskUpdatedEventFromJSONTyped(json: any, ignoreDiscriminator: boolean): SIPTaskUpdatedEvent {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'item': EnduroIngestSipTaskFromJSON(json['item']),
    };
}

export function SIPTaskUpdatedEventToJSON(value?: SIPTaskUpdatedEvent | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'item': EnduroIngestSipTaskToJSON(value.item),
    };
}

