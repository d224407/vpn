# TEMP-DISABLED: payment / subscription checks

Reason: paid server is currently failing, so subscription gating is disabled temporarily.
Search the codebase for `TEMP-DISABLED` to find every change.

| File | What was disabled | How to restore |
|------|-------------------|----------------|
| daemon/service/service_connect.go | Subscription expired / not-active check before connecting | Uncomment the block |
| daemon/service/preferences/account.go | MultiHop plan restriction (`IsCanConnectMultiHop`) | Delete the early `return nil` |
| ui/src/store/module-account.js | Expiry + free-trial banner messages | Remove the early `return null` (x2) |
| ui/src/components/settings/settings-account.vue | Subscription details block (plan, Upgrade, Add more time) | Remove `false &&` |
| ui/src/views/AccountLimit.vue | "Upgrade your subscription" button | Remove `false &&` |

Not changed: login/session flow, device limit logic, `PaymentMethod` field (data only).
