import * as client from 'typed-rest-client';

interface ITodo {
  ID: string;
  Text: string;
}

export async function GetTodo(): Promise<client.IRestResponse<ITodo[]>> {
  let rest: client.RestClient = new client.RestClient('rest-samples', "http://localhost:8080");
  let response: client.IRestResponse<ITodo[]> = await rest.get<ITodo[]>('todo');
  return response;
}
