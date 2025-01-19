import { Injectable } from '@nestjs/common';
import { PrismaService } from 'prisma/prisma.service';
import { Feedback, Prisma } from '@prisma/client';

@Injectable()
export class FeedbackService {
  constructor(private prisma: PrismaService) {}

  async createFeedback(data: Prisma.FeedbackCreateInput): Promise<Feedback> {
    return this.prisma.feedback.create({
      data,
    });
  }

  async getFeedbacks(): Promise<Feedback[]> {
    return this.prisma.feedback.findMany();
  }
}
