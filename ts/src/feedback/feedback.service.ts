import { Injectable, BadRequestException } from '@nestjs/common';
import { PrismaService } from 'prisma/prisma.service';
import {
  CreateFeedbackDto,
  FeebackReviewsDto,
  FeedbackResponseDto,
} from './feedback.dto';
import { Feedback } from '@prisma/client';

@Injectable()
export class FeedbackService {
  constructor(private prisma: PrismaService) {}

  async getUserReviews(reviewerId: string): Promise<FeebackReviewsDto[]> {
    const reviews = await this.prisma.review.findMany({
      where: {
        reviewers: {
          some: {
            reviewerId,
          },
        },
      },
      include: {
        employee: true,
      },
    });

    return reviews.map((review) => ({
      id: review.id,
      employeeName: review.employee.name,
      review: review.review,
    }));
  }

  async createFeedback(
    reviewerId: string,
    data: CreateFeedbackDto,
  ): Promise<FeedbackResponseDto> {
    // Check if the review assignment exists
    const assignment = await this.prisma.reviewAssignment.findFirst({
      where: {
        reviewId: data.reviewId,
        reviewerId,
      },
    });

    if (!assignment) {
      throw new BadRequestException('You are not assigned to this review.');
    }

    const feedback = await this.prisma.feedback.create({
      data: {
        reviewId: data.reviewId,
        reviewerId,
        feedback: data.feedback,
      },
    });

    // Remove the assignment after feedback is given
    await this.prisma.reviewAssignment.deleteMany({
      where: {
        reviewId: data.reviewId,
        reviewerId,
      },
    });

    return this.mapToFeedbackResponseDto(feedback);
  }

  async getFeedbacks(reviewerId: string): Promise<FeedbackResponseDto[]> {
    const feedbacks = await this.prisma.feedback.findMany({
      where: { reviewerId },
    });
    return feedbacks.map(this.mapToFeedbackResponseDto);
  }

  async getReviewFeedbacks(reviewId: string): Promise<FeedbackResponseDto[]> {
    const feedbacks = await this.prisma.feedback.findMany({
      where: { reviewId },
    });
    return feedbacks.map(this.mapToFeedbackResponseDto);
  }

  private mapToFeedbackResponseDto(feedback: Feedback): FeedbackResponseDto {
    return {
      id: feedback.id,
      reviewId: feedback.reviewId,
      reviewerId: feedback.reviewerId,
      feedback: feedback.feedback,
    };
  }
}
