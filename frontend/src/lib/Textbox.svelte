<script lang="ts">
    import { onMount } from "svelte";
    import { LoadNote, SaveNote } from "../../wailsjs/go/main/App";

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

<div>
    <textarea bind:value={content}></textarea>
    <button on:click={save}>Save</button>
    {#if status === "Saving..."}
        <p>Saving...</p>
    {:else if status === "Saved"}
        <p>Saved ✅</p>
    {/if}
</div>
