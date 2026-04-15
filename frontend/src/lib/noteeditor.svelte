<script lang="ts">
    import { onMount } from "svelte";
    import { LoadNote, SaveNote } from "../../wailsjs/go/main/App";
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

    onMount(async () => {
        content = await LoadNote();
        hasLoaded = true;
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
        content; // Placed here so when content changes it triggers effect
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
