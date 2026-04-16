<script lang="ts">
    import { onMount } from "svelte";
    import {
        LoadNote,
        SaveNote,
        ToggleAlwaysOnTop,
    } from "../../wailsjs/go/main/App";
    import { Button } from "$lib/components/ui/button";
    import { Textarea } from "$lib/components/ui/textarea";

    let content = $state("");
    let status = $state<"idle" | "typing" | "saving" | "saved" | "error">(
        "idle",
    );
    let hasLoaded = $state(false);
    let hasUserEdited = $state(false);
    let saveTimer: ReturnType<typeof setTimeout> | null = null;
    let statusTimer: ReturnType<typeof setTimeout> | null = null;

    onMount(() => {
        async function init() {
            try {
                content = await LoadNote();
                hasLoaded = true;
            } catch (err) {
                console.error(err);
                status = "error";
            }
        }

        init();

        window.addEventListener("keydown", handleKeydown);

        return () => {
            window.removeEventListener("keydown", handleKeydown);
            clearSaveTimer();
            clearStatusTimer();
        };
    });

    async function saveNow() {
        // Clear timers incase previous ones are still running
        clearSaveTimer();
        clearStatusTimer();

        status = "saving";

        try {
            await SaveNote(content);
            status = "saved";

            statusTimer = setTimeout(() => {
                status = "idle";
            }, 1500);
        } catch (err) {
            console.error(err);
            status = "error";
        }
    }

    $effect(() => {
        content; // track content changes for autosave
        if (!hasLoaded || !hasUserEdited) return;

        // clear timers
        clearSaveTimer();
        clearStatusTimer();

        // set status
        status = "typing";

        // start new timer
        saveTimer = setTimeout(saveNow, 1000);
    });

    function clearSaveTimer() {
        if (saveTimer) {
            clearTimeout(saveTimer);
            saveTimer = null;
        }
    }

    function clearStatusTimer() {
        if (statusTimer) {
            clearTimeout(statusTimer);
            statusTimer = null;
        }
    }

    function handleKeydown(event: KeyboardEvent) {
        const key = event.key.toLowerCase();
        const isMod = event.ctrlKey || event.metaKey;

        if (!isMod) return;

        switch (key) {
            case "s":
                if (status === "saving") return;
                event.preventDefault();
                saveNow();
                break;

            case "t":
                event.preventDefault();
                ToggleAlwaysOnTop();
                break;
        }
    }
</script>

<div class="h-screen flex flex-col p-4 gap-2">
    <Textarea
        class="flex-1 resize-none border rounded p-2"
        bind:value={content}
        oninput={() => {
            hasUserEdited = true;
        }}
    ></Textarea>

    <div class="flex justify-between">
        {#if status === "saving"}
            <p>Saving...</p>
        {:else if status === "saved"}
            <p>Saved ✅</p>
        {:else if status === "error"}
            <p>Failed to save ❌</p>
        {:else if status === "idle"}
            <p>Ready</p>
        {:else if status === "typing"}
            <p>Unsaved changes</p>
        {/if}

        <Button onclick={saveNow}>Save</Button>
    </div>
</div>
