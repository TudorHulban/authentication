# Ticket

## Ticket Types

There are several types of tickets.  
Ex. todo, sale, lead.  

## Ticket events

For each type of ticket there may be different events.

## Ticket status

Ticket status might be different than last event type. For example last event type could be internal note, but the customer should see just in work.

## Implementation high level

User groups would have levels. The level would be an empirical value that would allow to see real event info.  
Each user would imprint in the event a user group level which can be also lower than their level but not highr.  
For lower group levels the event seen would be last event at their level or below.  
Events could have dependent event, like for example internal note would add a in work event before.  

## Implementation low level

Ticket would not hold status.

Implementations:

```go
map[ticket kind]map[EventType]*TicketEventTypeInfo
```
