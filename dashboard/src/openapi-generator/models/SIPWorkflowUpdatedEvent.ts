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
import type { EnduroIngestSipWorkflow } from './EnduroIngestSipWorkflow';
import {
    EnduroIngestSipWorkflowFromJSON,
    EnduroIngestSipWorkflowFromJSONTyped,
    EnduroIngestSipWorkflowToJSON,
} from './EnduroIngestSipWorkflow';

/**
 * 
 * @export
 * @interface SIPWorkflowUpdatedEvent
 */
export interface SIPWorkflowUpdatedEvent {
    /**
     * Identifier of workflow
     * @type {number}
     * @memberof SIPWorkflowUpdatedEvent
     */
    id: number;
    /**
     * 
     * @type {EnduroIngestSipWorkflow}
     * @memberof SIPWorkflowUpdatedEvent
     */
    item: EnduroIngestSipWorkflow;
}

/**
 * Check if a given object implements the SIPWorkflowUpdatedEvent interface.
 */
export function instanceOfSIPWorkflowUpdatedEvent(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "id" in value;
    isInstance = isInstance && "item" in value;

    return isInstance;
}

export function SIPWorkflowUpdatedEventFromJSON(json: any): SIPWorkflowUpdatedEvent {
    return SIPWorkflowUpdatedEventFromJSONTyped(json, false);
}

export function SIPWorkflowUpdatedEventFromJSONTyped(json: any, ignoreDiscriminator: boolean): SIPWorkflowUpdatedEvent {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'item': EnduroIngestSipWorkflowFromJSON(json['item']),
    };
}

export function SIPWorkflowUpdatedEventToJSON(value?: SIPWorkflowUpdatedEvent | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'item': EnduroIngestSipWorkflowToJSON(value.item),
    };
}

