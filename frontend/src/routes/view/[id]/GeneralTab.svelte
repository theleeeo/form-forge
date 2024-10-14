<script lang="ts">
	import { formClient } from '$lib/formClient.js';
	import { Form, Question, TextQuestion } from '$lib/proto/form/v1/forms_pb.js';
	import { AccordionItem, Accordion, Card } from 'flowbite-svelte';
	import ViewForm from '$lib/components/view-form/ViewForm.svelte';

	export let form: Form;
	export let questions: Array<Question>;

	console.log(questions);
</script>

<ViewForm mode="view" {form} {questions} />
