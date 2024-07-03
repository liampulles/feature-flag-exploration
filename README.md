# Using bare Kubernetes for feature flags

Can we implement a basic form of feature flags without a special service? Yes!

If we limit the scope of feature flags to "config that can be changed dynamically at runtime, across apps", then
Kubernetes already provides for this. We just need to put the flags in a configmap, mount it as a volume in pods, and read
the mounted file whenever we query for flags in our app. Multiple apps can mount the same configmap, to enable cross-cutting
feature flag toggling.

There are disadvantages with this approach though:
- It doesn't help us for doing A/B testing, i.e. varying the flag
  based on user attributes, and tracking results. This is something most feature flag services allow for.
- There is a notable delay (around 1m15s in my setup) for the config to apply. And this delay is not tightly synchronized
  across apps (so flags might be intermittently read as on/off for a period across the broad system).
- It isn't usable by non-developers (unless they are comfortable with how to use kubernetes, or whatever
  deployment system deploys the configmap to kubernetes)

But, this is an example of how a bare Kubernetes setup can be more powerful than you might think.