import { Attribute, Event, Log } from "@cosmjs/stargate/build/logs"

export type RoadOperatorCreatedEvent = Event

export const getCreateRoadOperatorEvent = (
    log: Log,
): RoadOperatorCreatedEvent | undefined => {
    throw "Not implemented"
}

export const getCreatedRoadOperatorId = (
    createdRoadOperatorEvent: RoadOperatorCreatedEvent,
): string => {
    throw "Not implemented"
}
