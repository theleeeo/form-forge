import { writable } from 'svelte/store';

export type AlertType = {
	message: string;
	type: 'success' | 'warn' | 'error';
};

export const activeAlert = writable<AlertType | undefined>(undefined);

activeAlert.subscribe((alert) => {
	if (alert !== undefined) {
		setTimeout(() => {
			activeAlert.set(undefined);
		}, 5000);
	}
});
