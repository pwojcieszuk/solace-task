import { writable } from "svelte/store";

function createFavorites() {
    const { subscribe, set, update } = writable(
        new Map<number, { title: string; image: string }>(),
    );

    return {
        subscribe,
        add: (mal_id: number, title: string, image: string) =>
            update((n) => n.set(mal_id, { title: title, image: image })),
        remove: (mal_id: number) =>
            update((n) => {
                n.delete(mal_id);
                return n;
            }),
    };
}

export const favorites = createFavorites();
