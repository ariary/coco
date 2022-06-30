# 🥥 coco (c2c)

Command and server trial

3 component:
* **server**: deploy on attacker machine
* **agent**: deploy on target waiting for command from server
* **modules**: link to the Agent waiting for instruciton from agent

Communication:
* server <~> agent  (Websocket)
* agent <~> modules (IPC)


Agent can be:
* dynamically and OTA extended by **loading module**:
* statically built with **built-in modules**

The perk is that modules can be **custom at your convenience** (just follow a specific structure) and be **dynamically loaded** in the agent