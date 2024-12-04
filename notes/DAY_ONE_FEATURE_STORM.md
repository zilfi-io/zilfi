# What is dotenv feature set

- Share environment variables across teams
- Secrets management
  - AWS, encrypted mongo, maybe something else?
- Version history
  - Fully audited system
- Keep the environment variables in sync
  - Auto sync based on git hash
- Different types of environment variables
  - Some are shared (like 3rd party API keys)
  - Some just need to be present like X_API_KEY where it just needs to match in your local env
- Different environments
  - Local, dev, prod, etc.
- Easily swap to different environments
  - Can easily point local setup to staging for example
  - Should allow protection of who can point to what environment
- Permissions
  - Should allow protection of who can point to what environment
  - Git hooks to prevent check in?
  - Who can edit, add, delete environment variables
    - Fine grained access to specific environment variables or groups of environments
  - Trunk based system (for example)
    - Organizations, Teams, Projects, Environments, Users, Roles, likely others.
- Key rotation
  - Could be automatic based on:
    - Time
    - Changes in roster (e.g., someone leaves the team)
- Keys that need to be present but just need to be something
  - For example an x-api-key header for 2 services that you run locally that talk to each other
  - Allow the user to input and sync (across there apps locally) or just have the tool auto generate the key for them
- Some form of notification when env changes
  - Start with a change log stored in the cloud that says what changed and when
- CLI
  - Add
  - List
  - Edit
  - Remove
  - Rotate
  - Burn (mark a key as "do not use anymore")
  - Validate
  - Supports JSON and YAML

# Future features

- Secrets management with k8s intergration
  - Also argocd integration
- Depending on the environment you are in (e.g., test) certain .env vars are not
  required.
- Docker env vars vs. env vars in the terminal (e.g., nats//nats vs
  nats//localhost)
- In k8s discovery of endpoints is different depending on the namespace
  (nats//k8sservice...) Full URL vs. short url

# Anti-Patterns (Build Content Around These)
- Contingent env vars (i.e., if you use this OpenAI API key, use this assistant
  ID). Maybe this isn't something we solve for? Maybe this is a discovery
  mechanism, like look up the agent by name.

# Maybe features

- Allow securley sharing secrets so team memebers can use them locally without actually getting access
- Syncing deployed updates by restarting servers

## Git hash auto sync

Meaning that the tool is context aware that you shouldn't just always have "the most recent" environment variables, but you should have the appropriate environment variables for your point in the git history.
