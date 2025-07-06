import { Module } from '@nestjs/common';
import { APP_PIPE } from '@nestjs/core';
import { ValidationPipe } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { UserSettingsController } from './controllers/user-settings.controller';
import { UserSettingsService } from './services/user-settings.service';

@Module({
  imports: [],
  controllers: [AppController, UserSettingsController],
  providers: [
    AppService,
    UserSettingsService,
    {
      provide: APP_PIPE,
      useClass: ValidationPipe,
    },
  ],
})
export class AppModule {}
