 `Add new content at the top`
1. How do we do sharding?
2. Load balancer Design:
  - Application level (http)
  - Connection level (tcp/ip)
3. How to present or visualize the results of:
  - load balancing at different level
  - response to network partition
4. How to integrate personal project

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

## Action Items
| Items | Person Responsible  | Deadline |
| :---- | :------------------ | :---:|
| Prepare LB for individual function modules to allow kong configurations. | All | April 22 (Sat) |
| Start with individual work assigned. | All | April 22 (Sat) |
| Prepare action items for next meeting. | All | April 22 (Sat) |

### Decisions made
- Have inventory database sharded based on 2 store locations.
- App level will consist of basic user authentication.
- Order module will re-authenticate with User module for each request.


---