# Prompt Tracker

This code gets triggered with these Codex Hooks:

- UserPromptSubmit
- Stop

The users prompt is copied to an obsidan file with a diff of what was chagned for that prompt

## Obsidian Format

Each file in the project has a record with links to each modification that affected it next to a time stamp

For example:

### handler.go

_Modification #10/28/2026-10:54PM_

_etc etc_

--End of file--

### Modification #10/28/2026-10:54PM

_User's prompt_

_Diff of whats changed_

--End of file--
