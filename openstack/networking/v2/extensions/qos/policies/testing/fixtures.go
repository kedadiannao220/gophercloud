package testing

const GetPortResponse = `
{
    "port": {
        "id": "65c0ee9f-d634-4522-8954-51021b570b0d",
        "qos_policy_id": "591e0597-39a6-4665-8149-2111d8de9a08"
    }
}
`

const CreatePortRequest = `
{
    "port": {
        "network_id": "a87cc70a-3e15-4acf-8205-9b711a3531b7",
        "qos_policy_id": "591e0597-39a6-4665-8149-2111d8de9a08"
    }
}
`

const CreatePortResponse = `
{
    "port": {
        "network_id": "a87cc70a-3e15-4acf-8205-9b711a3531b7",
        "tenant_id": "d6700c0c9ffa4f1cb322cd4a1f3906fa",
        "id": "65c0ee9f-d634-4522-8954-51021b570b0d",
        "qos_policy_id": "591e0597-39a6-4665-8149-2111d8de9a08"
    }
}
`

const UpdatePortWithPolicyRequest = `
{
    "port": {
        "qos_policy_id": "591e0597-39a6-4665-8149-2111d8de9a08"
    }
}
`

const UpdatePortWithPolicyResponse = `
{
    "port": {
        "network_id": "a87cc70a-3e15-4acf-8205-9b711a3531b7",
        "tenant_id": "d6700c0c9ffa4f1cb322cd4a1f3906fa",
        "id": "65c0ee9f-d634-4522-8954-51021b570b0d",
        "qos_policy_id": "591e0597-39a6-4665-8149-2111d8de9a08"
    }
}
`

const UpdatePortWithoutPolicyRequest = `
{
    "port": {
        "qos_policy_id": null
    }
}
`

const UpdatePortWithoutPolicyResponse = `
{
    "port": {
        "network_id": "a87cc70a-3e15-4acf-8205-9b711a3531b7",
        "tenant_id": "d6700c0c9ffa4f1cb322cd4a1f3906fa",
        "id": "65c0ee9f-d634-4522-8954-51021b570b0d",
        "qos_policy_id": ""
    }
}
`
