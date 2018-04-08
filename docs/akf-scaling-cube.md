### X-Axis: Horizontal duplication

Each microservice of the Service Oriented Architecture (SOA) can be scaled independently from one another. These will be set up via an Auto-Scaling group in AWS such that it is allowed to grow or shink according to demand.

### Y-Axis: Functional Decomposition

The entire system will consist of a set of microservices that expose 
Solution: microservices
- Service 0: API Compositor (Kong)
  - Compose API interfaces of other microservices.
- Service 1: User Backend
  - User Account Information
  - User Preferences
  - Promotions (Optional)
- Service 2: Inventory / Location Backend
  - Starbucks Store Information
  - Starbucks Store Inventory
- Service 3: Ordering Backend
  - Online Order Placement
  - Online Order Fulfillment
- Service 4: Payment Backend
  - Payment Processing
  - Financial Reports

### Z-Axis: Data Partition

Multiple database clusters with application logics to distribute requests among the shards.
Not essential to all backends.

References:
1. http://microservices.io/articles/scalecube.html
