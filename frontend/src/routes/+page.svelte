<script lang="ts">
	import { formClient } from '$lib/formClient.js';
	import { Form, ListRequest } from '$lib/proto/form/v1/forms_pb.js';
	import FormPreview from './FormPreview.svelte';
	import { onMount } from 'svelte';

	let forms: Form[] = [];

	onMount(() => {
		console.log('Fetching forms');
		const request = new ListRequest({});
		formClient
			.list(request)
			.then((response) => {
				forms = response.forms;
			})
			.catch((error) => {
				console.error(error);
			});
	});
</script>

<div class="form-box grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
	{#each forms as form}
		<FormPreview {form} />
	{/each}
	<p>Hello</p>
</div>
