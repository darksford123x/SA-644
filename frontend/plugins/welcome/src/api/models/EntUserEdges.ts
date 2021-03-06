/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntDevice,
    EntDeviceFromJSON,
    EntDeviceFromJSONTyped,
    EntDeviceToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntUserEdges
 */
export interface EntUserEdges {
    /**
     * Device holds the value of the device edge.
     * @type {Array<EntDevice>}
     * @memberof EntUserEdges
     */
    device?: Array<EntDevice>;
}

export function EntUserEdgesFromJSON(json: any): EntUserEdges {
    return EntUserEdgesFromJSONTyped(json, false);
}

export function EntUserEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntUserEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'device': !exists(json, 'device') ? undefined : ((json['device'] as Array<any>).map(EntDeviceFromJSON)),
    };
}

export function EntUserEdgesToJSON(value?: EntUserEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'device': value.device === undefined ? undefined : ((value.device as Array<any>).map(EntDeviceToJSON)),
    };
}


