package runner

// EntrypointScript is a shell script that handles UID/GID mapping.
// It checks the ownership of the current directory (workspace) and creates a user
// with the same UID/GID if it doesn't exist.
const EntrypointScript = `#!/bin/sh
set -e

# Get the UID/GID of the current directory (workspace)
TARGET_UID=$(stat -c "%u" .)
TARGET_GID=$(stat -c "%g" .)

# If we are root, try to map to the target user
if [ "$(id -u)" = "0" ]; then
    # Check if a user with TARGET_UID already exists
    if ! getent passwd "$TARGET_UID" >/dev/null; then
        # Create group if it doesn't exist
        if ! getent group "$TARGET_GID" >/dev/null; then
            addgroup -g "$TARGET_GID" cm_group
        fi
        
        # Create user
        adduser -u "$TARGET_UID" -G "$(getent group "$TARGET_GID" | cut -d: -f1)" -D cm_user
    fi

    # Get the username for the UID
    USERNAME=$(getent passwd "$TARGET_UID" | cut -d: -f1)

    # Execute the command as the user
    # Try su-exec, then gosu, then su
    if command -v su-exec >/dev/null; then
        exec su-exec "$USERNAME" "$@"
    elif command -v gosu >/dev/null; then
        exec gosu "$USERNAME" "$@"
    else
        exec su "$USERNAME" -c "$*"
    fi
else
    # We are not root, just run the command
    exec "$@"
fi
`

// GetEntrypoint returns the entrypoint script content
func GetEntrypoint() string {
	return EntrypointScript
}
