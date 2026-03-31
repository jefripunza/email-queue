const testSendMessage = async () => {
  const response = await fetch("http://localhost:3000/api/message/send", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "X-API-Key": "test-key",
    },
    body: JSON.stringify({
      to: "jefripunza@gmail.com",
      subject: "Test Email",
      body: "This is a test email",
    }),
  });
  console.log(await response.json());
};
// testSendMessage();

const checkStatusMessage = async (key: string) => {
  const response = await fetch(
    `http://localhost:3000/api/message/status/${key}`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "X-API-Key": "test-key",
      },
    },
  );
  console.log(await response.json());
};
// checkStatusMessage("dfc9e68f-9faf-4d21-9ebc-84b92d7183fd");
