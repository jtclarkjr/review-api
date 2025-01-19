/*
  Warnings:

  - A unique constraint covering the columns `[reviewId,reviewerId]` on the table `Feedback` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[reviewId,reviewerId]` on the table `ReviewAssignment` will be added. If there are existing duplicate values, this will fail.

*/
-- CreateIndex
CREATE UNIQUE INDEX "Feedback_reviewId_reviewerId_key" ON "Feedback"("reviewId", "reviewerId");

-- CreateIndex
CREATE UNIQUE INDEX "ReviewAssignment_reviewId_reviewerId_key" ON "ReviewAssignment"("reviewId", "reviewerId");
