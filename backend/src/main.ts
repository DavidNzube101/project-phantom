import { NestFactory } from '@nestjs/core';
import { SwaggerModule, DocumentBuilder } from '@nestjs/swagger';
import { AppModule } from './app.module';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  // Swagger configuration
  const config = new DocumentBuilder()
    .setTitle('Project Phantom API')
    .setDescription(`
      Privacy-first browser backend API for user settings management.
      
      ## Privacy Features
      - Users are identified only by wallet address (no personal data stored)
      - Flexible settings schema allows for future privacy features
      - Minimal data collection - only essential user preferences
      - No tracking, analytics, or user behavior monitoring
      
      ## Authentication
      - No traditional authentication required
      - Users are identified by their wallet address
      - Settings are created on first save request
      
      ## Data Storage
      - All data stored in Firebase Firestore
      - Settings are flexible JSON objects
      - No browsing history or personal data stored
    `)
    .setVersion('1.0')
    .addTag('User Settings', 'Endpoints for managing user preferences and settings')
    .addServer('http://localhost:3000', 'Development server')
    .addServer('https://api.project-phantom.com', 'Production server')
    .build();

  const document = SwaggerModule.createDocument(app, config);
  SwaggerModule.setup('api', app, document, {
    swaggerOptions: {
      persistAuthorization: true,
      displayRequestDuration: true,
      filter: true,
      showExtensions: true,
      showCommonExtensions: true,
    },
    customSiteTitle: 'Project Phantom API Documentation',
    customCss: `
      .swagger-ui .topbar { display: none }
      .swagger-ui .info .title { color: #2c3e50; }
      .swagger-ui .info .description { color: #34495e; }
    `,
  });

  await app.listen(process.env.PORT ?? 3000);
  
  console.log(`🚀 Application is running on: http://localhost:${process.env.PORT ?? 3000}`);
  console.log(`📚 Swagger documentation available at: http://localhost:${process.env.PORT ?? 3000}/api`);
}
bootstrap();
