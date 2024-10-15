import { formClient } from '$lib/formClient.js';
import { Form, GetByIdRequest, GetQuestionsRequest } from '$lib/proto/form/v1/forms_pb.js';
import { ConnectError } from '@connectrpc/connect';

export const load = async ({ params }) => {
	const formResp = formClient.getById(
		new GetByIdRequest({
			baseId: params.id
		})
	);
	const questionResp = formClient.getQuestions(
		new GetQuestionsRequest({
			baseId: params.id
		})
	);

	const [form, questions] = await Promise.all([formResp, questionResp]);

	if (!form.form) {
		throw new Error('no form returned');
	}

	return {
		form: form.form,
		questions: questions.questions
	};
};
