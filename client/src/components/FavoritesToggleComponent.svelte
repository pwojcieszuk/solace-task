<script lang="ts">
    import { favorites } from "../stores";

    export let title: string;
    export let mal_id: number;
    export let image: string;

    function handleMousedown() {
        $favorites.has(mal_id)
            ? favorites.remove(mal_id)
            : favorites.add(mal_id, title, image);
    }
</script>

<div
    class="star opacity-80 hover:opacity-100"
    class:favorite={$favorites.has(mal_id)}
    title={`${$favorites.has(mal_id) ? "Remove from" : "Add to"} favorites`}
    on:mousedown={handleMousedown}
/>

<style>
    .star {
        position: absolute;
        z-index: 1;
        right: 1rem;
        top: 1rem;
        width: 3rem;
        height: 3rem;
        background-color: grey;

        clip-path: polygon(
            50% 0%,
            61% 35%,
            98% 35%,
            68% 57%,
            79% 91%,
            50% 70%,
            21% 91%,
            32% 57%,
            2% 35%,
            39% 35%
        );
    }

    .star:hover {
        cursor: pointer;
        background-color: #dc2626;
    }

    .star.favorite {
        background-color: #dc2626;
    }

    .star.favorite:hover {
        background-color: grey;
    }
</style>
