export interface AiServerRequest {
	queryInput: {
		text: {
			text: string;
			languageCode: "id";
		};
	};
}

export interface AiServerResponse {
	responseId: string;
	queryResult: {
		queryText: string;
		parameters: {
			title: string;
			"date-time": { date_time: string }[];
			description: string[];
			id?: string;
		};
		fulfillmentText: string;
		intent: {
			displayName: string;
		};
		allRequiredParamsPresent: boolean;
	};
}
