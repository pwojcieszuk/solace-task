<script lang="ts">
    import type { PageData } from "./$types";
    import AnimeListItemComponent from "../components/AnimeListItemComponent.svelte";

    export let data: PageData;
</script>

{#each data.recommended.data as recommendation}
    <AnimeListItemComponent {recommendation} />
{/each}
