
---

## Meeting - 01

| LOCATION | ONLINE |
|:----|:----|
| **Date** | 2018/04/04 |
| **Facilitator** | N/A |
| **Timer** | N/A |
| **Note Taker** | Bo |
| **Attendees** | Brian Bamsch<br>Bo Liu<br>Will Maynard<br>Masi Nazarian<br>Vimmi Swami<br> |

### Agenda Topics
1. Project commencement
2. Discuss and decide on which application to work on for the team project

## Action Items
| Items | Person Responsible  | Deadline |
| :---- | :------------------ | :---:|
| Prepare description of at least 3 system components for discussion on Saturday. | All | April 7 (Sat) |
| Propose at least one system architecture or design pattern. Ref: [The Practice of Cloud System Administration ](http://ptgmedia.pearsoncmg.com/images/9780321943187/samplepages/9780321943187.pdf) | All | April 7 (Sat) |
| Study the personal project description and think about how to integrate it into the team project. | All | April 7 (Sat) |
| Read `high-level-design-doc.md` for other TODOs. | All | April 7 (Sat) |

### Decisions made
- Use Starbucks as the application for the team project.


---

## Meeting - 02

| LOCATION | ONLINE |
|:----|:----|
| **Date** | 2018/04/15 |
| **Facilitator** | N/A |
| **Timer** | N/A |
| **Note Taker** | Vimmi |
| **Attendees** | Brian Bamsch<br>Bo Liu<br>Masi Nazarian<br>Vimmi Swami<br> |

### Agenda Topics
1. Discuss Sharding component in project
2. Discuss and decide on component level design and implementation
3. Discuss and assign work to team members:
- Bo : User module
- Vimmi : Order module
- Masi and Brian : Kong setup and Inventory module
- Brian : Payment module  

## Action Items
| Items | Person Responsible  | Deadline |
| :---- | :------------------ | :---:|
| Prepare LB for individual function modules to allow kong configurations. | All | April 21 (Sat) |
| Start with individual work assigned. | All | April 21 (Sat) |
| Prepare action items for next meeting. | All | April 21 (Sat) |

### Decisions made
- Have inventory database sharded based on 2 store locations.
- App level will consist of basic user authentication.
- Order module will re-authenticate with User module for each request.


---
## Meeting - 03

| LOCATION | ONLINE |
|:----|:----|
| **Date** | 2018/04/20 |
| **Facilitator** | N/A |
| **Timer** | N/A |
| **Note Taker** | Vimmi |
| **Attendees** | Brian Bamsch<br>Bo Liu<br>Masi Nazarian<br>Vimmi Swami<br> |

### Agenda Topics
1. Discuss Authentication of user and its processing
2. Discuss Order and payment Interface
3. Discuss TODO items for meeting on 4/21/2018.


## Action Items
| Items | Person Responsible  | Deadline |
| :---- | :------------------ | :---:|
| Get basic API interface setup with backend DB | All | April 21 (Sat) |

### Decisions made
- Authenticating token Id will be sent to GO backend servers too by frontend app server.
- Each bckend server will have sufficient data to decode token to user_id without communicating with User backend.
- Each Backend server will send auth token id along with its request to other backend APIs while communicating.

---

## Meeting - 04

| LOCATION | Library + Whatsapp Group |
|:----|:----|
| **Date** | 2018/04/26 |
| **Facilitator** | N/A |
| **Timer** | N/A |
| **Note Taker** | Bo |
| **Attendees** | Brian Bamsch<br>Bo Liu<br>Vimmi Swami<br> |

### Agenda Topics
1. Deploy Orders API to EC2
2. Update Kong config
3. Testing Orders and User API

### Decision
Meeting in person on Sunday morning 9:00 am for integration tests.

---

## Meeting - 05

| LOCATION | Library + Whatsapp Group |
|:----|:----|
| **Date** | 2018/04/29 |
| **Facilitator** | N/A |
| **Timer** | N/A |
| **Note Taker** | Vimmi |
| **Attendees** | Brian Bamsch<br>Bo Liu<br>Vimmi Swami<br> |

### Agenda Topics
1. Have basic integration done of User, Order, Inventory and Payment APIs
2. Update Kong config to remove duplicate Order references
3. Testing UI for end to end working

### Decision
Have Auto scaling done for individual modules and meet again for proper testing and issues.

---

## Meeting - 06

| LOCATION | Library + Whatsapp Group |
|:----|:----|
| **Date** | 2018/05/02 |
| **Facilitator** | N/A |
| **Timer** | N/A |
| **Note Taker** | Bo |
| **Attendees** | Brian Bamsch<br>Bo Liu<br>Vimmi Swami<br> |

### Agenda Topics
1. Debugging EC2 related issue: cannot connect to Cassandra from Go API.
2. Discussed how to setup auto-scaling group and load balancer.
3. Performed end to end integration tests.

### Action Items:
1. Fix the authentication issue caused by local storage. (Bo)
2. Sharding inventory database based on store id. (Brian)
3. Investigate error message after deleting order. (Vimmi)

---

## Meeting - 07

| LOCATION | Zoom + Whatsapp Group |
|:----|:----|
| **Date** | 2018/05/04 |
| **Facilitator** | N/A |
| **Timer** | N/A |
| **Note Taker** | Vimmi |
| **Attendees** | Brian Bamsch<br>Bo Liu<br>Vimmi Swami<br> |

### Agenda Topics
1. Integration testing.
2. Resolved some frontend related issues.
3. Discuss and prepare report for assignment submission.

### Action Items:
1. Fix the delete javascript related issue. (Vimmi)
2. Fix the http client related "many open file descriptors issue". (Vimmi)
3. Pass possible integration testing cases.
4. Submmit project report. (ALL)
