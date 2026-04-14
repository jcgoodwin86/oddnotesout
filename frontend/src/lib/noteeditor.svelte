<script lang="ts">
    import { onMount } from "svelte";
    import { LoadNote, SaveNote } from "../../wailsjs/go/main/App";
    import { Button } from "$lib/components/ui/button";
    import { Textarea } from "$lib/components/ui/textarea";

    let content: string = "";
    let status: "idle" | "saving" | "saved" | "error" = "idle";

    onMount(async () => {
        content = await LoadNote();
    });

    let save = async () => {
        status = "saving";
        try {
            await SaveNote(content);
            status = "saved";

            setTimeout(() => {
                status = "idle";
            }, 1500);
        } catch (err) {
            console.error(err);
            status = "error";
        }
    };
</script>

<div class="h-screen flex flex-col p-4 gap-2">
    <Textarea class="flex-1 resize-none border rounded p-2" bind:value={content}
    ></Textarea>

    <div class="flex justify-between">
        {#if status === "saving"}
            <p>Saving...</p>
        {:else if status === "saved"}
            <p>Saved ✅</p>
        {:else if status === "error"}
            <p>Failed to save ❌</p>
        {:else if status === "idle"}
            <p>idle</p>
        {/if}

        <Button onclick={save}>Save</Button>
    </div>
</div>
