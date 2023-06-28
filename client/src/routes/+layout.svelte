<script lang="ts">
    import "../app.css";
    import type { LayoutData } from "./$types";
    import FavoritesComponent from "../components/FavoritesComponent.svelte";
    export let data: LayoutData;
</script>

<div class="bg-hive-pattern bg-cover h-screen overflow-hidden">
    <div class="overflow-y-auto h-screen flex">
        {#if data.favorites.size !== 0}
            <div class="min-w-fit">
                <FavoritesComponent favorites={data.favorites} />
            </div>
        {/if}
        <div class="flex flex-wrap justify-center"><slot /></div>
    </div>
</div>
