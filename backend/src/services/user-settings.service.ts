import { Injectable, NotFoundException } from '@nestjs/common';
import { getFirestore } from '../config/firebase.config';
import { SaveUserSettingsDto, GetUserSettingsDto, UserSettingsResponseDto } from '../dto/user-settings.dto';

@Injectable()
export class UserSettingsService {
  private readonly db = getFirestore();
  private readonly collectionName = 'user_settings';

  async saveUserSettings(saveUserSettingsDto: SaveUserSettingsDto): Promise<UserSettingsResponseDto> {
    const { walletAddress, settings } = saveUserSettingsDto;
    
    const now = new Date();
    const userSettingsData = {
      walletAddress,
      settings: settings || {},
      createdAt: now,
      updatedAt: now,
    };

    // Check if user already exists
    const existingUser = await this.getUserSettingsByWalletAddress(walletAddress);
    
    if (existingUser) {
      // Update existing user settings
      const updatedData = {
        ...userSettingsData,
        createdAt: existingUser.createdAt, // Preserve original creation date
      };
      
      await this.db
        .collection(this.collectionName)
        .doc(walletAddress)
        .update(updatedData);
      
      return updatedData;
    } else {
      // Create new user settings
      await this.db
        .collection(this.collectionName)
        .doc(walletAddress)
        .set(userSettingsData);
      
      return userSettingsData;
    }
  }

  async getUserSettings(getUserSettingsDto: GetUserSettingsDto): Promise<UserSettingsResponseDto> {
    const { walletAddress } = getUserSettingsDto;
    
    const userSettings = await this.getUserSettingsByWalletAddress(walletAddress);
    
    if (!userSettings) {
      throw new NotFoundException(`User settings not found for wallet address: ${walletAddress}`);
    }
    
    return userSettings;
  }

  async getUserSettingsByWalletAddress(walletAddress: string): Promise<UserSettingsResponseDto | null> {
    const doc = await this.db
      .collection(this.collectionName)
      .doc(walletAddress)
      .get();

    if (!doc.exists) {
      return null;
    }

    const data = doc.data();
    if (!data) {
      return null;
    }

    return {
      walletAddress: data.walletAddress,
      settings: data.settings,
      createdAt: data.createdAt.toDate(),
      updatedAt: data.updatedAt.toDate(),
    };
  }

  async updateUserSettings(walletAddress: string, settings: Record<string, any>): Promise<UserSettingsResponseDto> {
    const existingUser = await this.getUserSettingsByWalletAddress(walletAddress);
    
    if (!existingUser) {
      throw new NotFoundException(`User settings not found for wallet address: ${walletAddress}`);
    }

    const updatedData = {
      ...existingUser,
      settings: { ...existingUser.settings, ...settings },
      updatedAt: new Date(),
    };

    await this.db
      .collection(this.collectionName)
      .doc(walletAddress)
      .update(updatedData);

    return updatedData;
  }

  async deleteUserSettings(walletAddress: string): Promise<void> {
    const existingUser = await this.getUserSettingsByWalletAddress(walletAddress);
    
    if (!existingUser) {
      throw new NotFoundException(`User settings not found for wallet address: ${walletAddress}`);
    }

    await this.db
      .collection(this.collectionName)
      .doc(walletAddress)
      .delete();
  }
} 