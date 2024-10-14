// import { ListResponse } from '$lib/proto/form/v1/forms';
// import { formClient } from '$lib/server/form_client';
// import { CreateRequest, CreateResponse } from '$lib/proto/form/v1/forms';

// export async function load() {
// 	try {
// 		const response = await new Promise<ListResponse>((resolve, reject) => {
// 			formClient.list({}, (error, response) => {
// 				if (error) {
// 					reject(error);
// 				} else {
// 					resolve(response);
// 				}
// 			});
// 		});

// 		return {
// 			forms: response.forms
// 		};
// 	} catch (error) {
// 		console.error('An error occurred while loading forms:', error);
// 		return {
// 			forms: []
// 		};
// 	}
// }

// export async function createForm(data: CreateRequest): Promise<CreateResponse> {
// 	try {
// 		const response = await new Promise<CreateResponse>((resolve, reject) => {
// 			formClient.create(data, (error, response) => {
// 				if (error) {
// 					reject(error);
// 				} else {
// 					resolve(response);
// 				}
// 			});
// 		});

// 		return response;
// 	} catch (error) {
// 		console.error('An error occurred while creating the form:', error);
// 		throw error;
// 	}
// }

// export const actions = {
// 	default: async () => {
// 		try {
// 			const response = await new Promise<ListResponse>((resolve, reject) => {
// 				formClient.list({}, (error, response) => {
// 					if (error) {
// 						reject(error);
// 					} else {
// 						resolve(response);
// 					}
// 				});
// 			});

// 			return {
// 				forms: response.forms
// 			};
// 		} catch (error) {
// 			console.error('An error occurred while loading forms:', error);
// 			return {
// 				forms: []
// 			};
// 		}
// 	}
// };
