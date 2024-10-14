// import { ListResponse } from "$lib/proto/form/v1/forms";
// import { formClient } from "$lib/server/form_client";

// export async function load() {
//     try {
//         const response = await new Promise<ListResponse>((resolve, reject) => {
//             formClient.list({}, (error, response) => {
//                 if (error) {
//                     reject(error);
//                 } else {
//                     resolve(response);
//                 }
//             });
//         });

//         return {
//             forms: response.forms
//         };
//     } catch (error) {
//         console.error("An error occurred while loading forms:", error);
//         return {
//             forms: []
//         };
//     }
// }
//     const response = await fetch('http://localhost:8899/form.v1.FormService/List', {
//         method: 'POST',
//         mode: 'no-cors',
//         headers: {
//             'Content-Type': 'application/json',
//         },
//         body: JSON.stringify({
//             limit: 10,
//         }),
//     });

//     const parsedData = ListResponse.fromJSON(await response.json())

//     return {
//         forms: parsedData.forms
//     };
// } catch (error) {
//     return {
//         forms: []
//     };
// }w
