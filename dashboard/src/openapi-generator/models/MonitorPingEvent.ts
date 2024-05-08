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
 * @interface MonitorPingEvent
 */
export interface MonitorPingEvent {
    /**
     * 
     * @type {string}
     * @memberof MonitorPingEvent
     */
    message?: string;
}

/**
 * Check if a given object implements the MonitorPingEvent interface.
 */
export function instanceOfMonitorPingEvent(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function MonitorPingEventFromJSON(json: any): MonitorPingEvent {
    return MonitorPingEventFromJSONTyped(json, false);
}

export function MonitorPingEventFromJSONTyped(json: any, ignoreDiscriminator: boolean): MonitorPingEvent {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'message': !exists(json, 'message') ? undefined : json['message'],
    };
}

export function MonitorPingEventToJSON(value?: MonitorPingEvent | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'message': value.message,
    };
}

