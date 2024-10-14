// place files you want to import through the `$lib` alias in this folder.

// import { setContext } from 'svelte';
import { createConnectTransport } from '@connectrpc/connect-web';
import { createPromiseClient } from '@connectrpc/connect';
import { FormService } from '$lib/proto/form/v1/forms_connect.js';

const transport = createConnectTransport({
	baseUrl: 'http://localhost:8899'
});
export const formClient = createPromiseClient(FormService, transport);

// setContext('transport', transport);
