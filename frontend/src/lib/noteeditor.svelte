<script lang="ts">
    import { onMount } from "svelte";
    import { LoadNote, SaveNote } from "../../wailsjs/go/main/App";
    import { Button } from "$lib/components/ui/button";
    import { Textarea } from "$lib/components/ui/textarea";

    let content: string = "";
    let status: string = "Idle";

    onMount(async () => {
        content = await LoadNote();
    });

    let save = async () => {
        status = "Saving...";
        await SaveNote(content);
        status = "Saved";
    };
</script>

<div class="h-screen flex flex-col p-4 gap-2">
    <Textarea class="flex-1 resize-none border rounded p-2" bind:value={content}
    ></Textarea>

    <div class="flex justify-between">
        {#if status === "Saving..."}
            <p>Saving...</p>
        {:else if status === "Saved"}
            <p>Saved ✅</p>
        {/if}

        <Button onclick={save}>Save</Button>
    </div>
</div>
