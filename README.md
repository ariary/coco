
<div align=center>
<h1>🥥 coco (c2c)</h1>
<img src=./img/coco.png widtH=30%>
<h3><strong>Dynamically extendable Command and server trial</strong></h3>
</div>


## Usage

* [Launch server and agent](./c2c/README.md)
* [Write modules](./modules/README.md)

## Overview
***TL;DR***:The perk is that modules can be **custom at your convenience** (just follow a specific structure) and be **dynamically loaded** in the agent

3 components:
* **server**: deploy on attacker machine
* **agent**: deploy on target waiting for command from server
* **modules**: link to the Agent waiting for instruciton from agent

Communication:
* server <~> agent  (Websocket)
* agent <~> modules (IPC)


Agent can be:
* dynamically and OTA extended by **loading module**:
* statically built with **built-in modules**

