# Inventory
The inventory contains information about our agents. The concept is very similar to Ansible's [inventory](https://docs.ansible.com/ansible/latest/inventory_guide/intro_inventory.html) file.

```yaml title="inventory.yaml"
nodes:
 - name: "lab1"
   location: "192.168.0.99:8080"
   key: "Z6wC4goD7V2EiL4XuecTuo8jVxfvwVxs"

 - name: "lab2"
   location: "192.168.0.100:8080"
   key: "EcwxaMO3kyBaKETesxInx7ga3Ti93gvI"

groups:
 - name: "labs"
   nodes: ["lab1", "lab2"]
```

## Nodes
The nodes field is an array of nodes (or servers, both terms are used interchangeably) where the agents are running. It consists of the following items:

- **name** is the custom name you want to give to a node.
- **location** is the URL where the agent is running.
- **key** is the Agent's API key.

## Groups
The groups field is an array of groups, where you can group together multiple nodes to avoid referring to them individually. It consists of the following items:

- **name** is the custom name you want to give to a group.
- **nodes** is a list of nodes you want to add to a group.