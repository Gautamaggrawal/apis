from django.db import models
import datetime


class Userdata(models.Model):
    username = models.CharField(db_column='userName', max_length=25)
    emailid = models.CharField(
        db_column='emailId',
        primary_key=True,
        max_length=50)
    phoneno = models.CharField(db_column='phoneNo', max_length=10)
    password = models.CharField(max_length=50)
    datetime = models.DateTimeField(
        db_column='dateTime',
        blank=True,
        null=True,
        auto_now=True)

    class Meta:
        db_table = 'userData'
