declare global {
  type KeyPair = {
    Private: string;
    Public: string;
  };
  interface Window {
    backend: {
      genKeyPair(): Promise<KeyPair>;
      add(title: string, uuid: string): Promise<{ requestKeys: string[] }>;
      save(id: string, text: string): Promise<{ requestKeys: string[] }>;
      toggle(id: string): Promise<{ requestKeys: string[] }>;

      toggleAll(id: string[]): Promise<{ requestKeys: string[] }>;
      clearCompleted(id: string[]): Promise<{ requestKeys: string[] }>;

      destroy(id: string): Promise<{ requestKeys: string[] }>;
      getTodos(): Promise<{
        logs: string;
        reqKey: string;
        result: {
          data: Array<{
            completed: boolean;
            deleted: boolean;
            id: string;
            title: string;
          }>;
          status: "success";
        };
      }>;
    };
  }
}

export {};
