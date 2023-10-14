openapi: '3.0.0'
info:
  version: '1.0.0'
  title: 'PromptShield - PII and Prompt Injection Verification Service'
  description: "PromptShield's API provides a mechanism to strip PII and check for prompt injection, ensuring safe text processing."
  contact:
    name: 'Support'
    email: 'admin@safetorun.com'

servers:
  - url: 'https://prompt.safetorun.com'
    description: 'Production server'

paths:
  /keep:
    post:
      x-amazon-apigateway-integration:
        uri: ${lambda_keep_arn}
        passthroughBehavior: "when_no_match"
        httpMethod: "POST"
        type: "aws_proxy"
      summary: 'Verify and Analyze Prompt'
      description: 'This endpoint accepts a text prompt, strips PII, and checks it for prompt injection, returning an injection score.'
      operationId: 'verifyPrompt'
      security:
        - ApiKeyAuth: [ ]
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/PromptRequest'
      responses:
        '200':
          description: 'Successful operation.'
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/InjectionScoreResponse'
        '400':
          description: 'Bad request. The prompt field is missing or invalid.'
        '500':
          description: 'Internal server error.'
  /moat:
    post:
      x-amazon-apigateway-integration:
        uri: ${lambda_moat_arn}
        passthroughBehavior: "when_no_match"
        httpMethod: "POST"
        type: "aws_proxy"
      summary: 'Get your ring 3 shield prompt'
      description: 'This endpoint accepts a text prompt, strips PII, and checks it for prompt injection, returning an injection score.'
      operationId: 'buildShield'
      security:
        - ApiKeyAuth: [ ]
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/PromptShieldRequest'
      responses:
        '200':
          description: 'Successful operation.'
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/PromptShieldResponse'
        '400':
          description: 'Bad request. The prompt field is missing or invalid.'
        '500':
          description: 'Internal server error.'
components:
  securitySchemes:
    ApiKeyAuth:
      type: 'apiKey'
      in: 'header'
      name: 'x-api-key'
      description: 'API key required for AWS API Gateway'
  schemas:
    PromptRequest:
      type: 'object'
      required:
        - 'prompt'
      properties:
        prompt:
          type: 'string'
          description: 'The text prompt to be verified.'

    PromptShieldRequest:
      type: 'object'
      required:
        - 'prompt'
      properties:
        prompt:
          type: 'string'
          description: 'The text prompt to be verified.'

    PromptShieldResponse:
      type: 'object'
      properties:
        shielded_prompt:
          type: 'string'
          description: 'The shielded prompt.'

    InjectionScoreResponse:
      type: 'object'
      properties:
        injection_score:
          type: 'number'
          format: 'float'
          description: 'The score indicating the likelihood of prompt injection.'
