<script lang="ts">
    import "../app.css";
    import FavoritesComponent from "../components/FavoritesComponent.svelte";
    import { favorites } from "../stores";

    export let data: {
        favorites: Map<number, { title: string; image: string }>;
    };
    favorites.set(data.favorites);
</script>

<div class="bg-hive-pattern bg-cover h-screen overflow-hidden">
    <div class="overflow-y-auto h-screen flex">
        {#if $favorites.size !== 0}
            <div class="min-w-fit">
                <FavoritesComponent favorites={$favorites} />
            </div>
        {/if}
        <div class="flex flex-wrap justify-center"><slot /></div>
    </div>
</div>
