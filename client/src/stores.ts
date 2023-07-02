import { writable } from "svelte/store";
import { PUBLIC_CLIENT_API_URL } from "$env/static/public";

const apiUrl = PUBLIC_CLIENT_API_URL;

function createFavorites() {
    const { subscribe, set, update } = writable(
        new Map<number, { title: string; image: string }>(),
    );

    const add = async (mal_id: number, title: string, image: string) => {
        await fetch(`${apiUrl}/favorites`, {
            method: "POST",
            body: JSON.stringify({ ID: mal_id, title: title, image: image }),
            headers: new Headers({
                "Content-Type": "application/json",
            }),
        });
        update((n) => n.set(mal_id, { title: title, image: image }));
    };

    const remove = async (mal_id: number) => {
        await fetch(`${apiUrl}/favorites/${mal_id}`, {
            method: "DELETE",
        });
        update((n) => {
            n.delete(mal_id);
            return n;
        });
    };

    return {
        subscribe,
        set,
        add,
        remove,
    };
}

export const favorites = createFavorites();
