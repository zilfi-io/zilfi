Diff Env Files
- [ ] Does my .env file differ from the one in main (.env.example or .env.sample)?
  - That is, do they have the same keys (not necessarily the same values)?
- [ ] Add this to your .zshrc and if there is a .envy.json check for a .env /
  .env.example mismatch, it will tell you.
  - This should run on:
    - entering the directory
    - saving a .env file
- [ ] There should be a .envy.json file next to the .git directory.
  - [ ] Lower level directories can also have a .envy.json file to override
    the parent directory.
- [ ] There should also be a system level .envy.config file that can be
  overridden by the project level .envy.json file.

What is a .envy.json file?

{
    "locations": [
        {
            "path": "cmd",
            "sample": ".env.example",
            "environments": [
                {
                    "name": "local",
                    "file": ".env.local",
                    "sample": nil,
                    "updateBehavior": "auto"
                },
                {
                    "name": "dev",
                    "file": ".env.dev",
                    "updateBehavior": "manual"
                },
                {
                    "name": "prod",
                    "file": ".env.prod",
                    "sample": ".env.prod.sample",
                    "updateBehavior": "ask"
                }
            ],
            "updateBehavior": "auto"
        }, {
            "path": "cmd/server",
            "sample": ".env.example",
            "environments": [
                {
                    "name": "local",
                    "sample": ".env.local",
                    "updateBehavior": "auto"
                },
                {
                    "name": "dev",
                    "sample": ".env.dev",
                    "updateBehavior": "manual"
                },
                {
                    "name": "prod",
                    "sample": ".env.prod",
                    "updateBehavior": "manual"
                }
            ],
        }
    ]
}
