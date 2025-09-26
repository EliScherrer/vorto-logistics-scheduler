# vorto-logistics-scheduler



The application is a simple work scheduling platform in the logistics industry. The primary users

are drivers who log into the application to receive work from the server. The unit of work is a

load: an order to pick up a payload from one location and drop it off at another. Loads are

submitted to the system via an admin, who has a simple UI to submit loads and track their

progress. The system then dispatches loads to the drivers efficiently, who in turn inform the

system when each stop (i.e. the pickup and dropoff) is completed.

The specific requirements of this application are as follows:



Driver Interface:

- A driver can be on shift, during which time he can be assigned loads; or off shift, during which time no work will be assigned to him While not on shift…

- Allow a driver to start a shift, at which time they provide their current location as a (latitude, longitude) pair

- Shifts are indefinite in duration; they do not end until the driver voluntarily ends them

- While on shift, a driver may be assigned loads from the server. When assigned a load…

- Show the load’s pickup and dropoff locations on a map

- Provide a button to complete the next task, which is either to pickup or

dropoff the payload

- Completing each stop should inform the system that the driver is

now at that location

- Upon completing the dropoff, the load itself is completed and

should be removed from view

- If the pickup has not yet been completed, provide a button to reject the

load, which also ends the driver’s shift

- If a driver is on shift and has no load assigned, provide a button to end shift



Admin Interface:

- Provide a tool for submitting loads to the application

- A load must simply have a pickup location and a dropoff location, both in the form of a (latitude, longitude)

- The tool must provide a map interface so that shippers can simply click on two locations and submit

- Provide a list of all loads in the system, with their current statuses

- Possible statuses are awaiting driver, in progress, and completed.



Backend requirements

- All drivers, shifts, and loads should be persisted to a database

- Load assignment should be efficient. E.g. all else equal, loads should be dispatched to minimize driven miles



You do not need to worry about user authentication. Assume that drivers simply provide a user

name when beginning a session; a driver is created if one does not exist. Drivers should be able

to log out while on shift or with a load, and log back in still on the same shift and load. That is,

logging out doesn’t end a shift or unassign a load.



You may use any language and technology to build this application. The driver and admin

interfaces should be web-based - not mobile - though these users should log in via separate

portals. Your final submission should be emailed in a zip, with all files and instructions needed to

install and run the program, preferably containerized.