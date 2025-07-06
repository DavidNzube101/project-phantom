import { Module } from '@nestjs/common';
import { APP_PIPE } from '@nestjs/core';
import { ValidationPipe } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { UserSettingsController } from './controllers/user-settings.controller';
import { UserSettingsService } from './services/user-settings.service';
import { UserSettingsMockService } from './services/user-settings-mock.service';

@Module({
  imports: [],
  controllers: [AppController, UserSettingsController],
  providers: [
    AppService,
    // Use mock service for development (comment out to use Firebase)
    UserSettingsMockService,
    // UserSettingsService, // Uncomment to use Firebase
    {
      provide: APP_PIPE,
      useClass: ValidationPipe,
    },
  ],
})
export class AppModule {}
