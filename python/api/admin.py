from django.contrib import admin

# Register your models here.
from api.models import Userdata

class UserdataAdmin(admin.ModelAdmin):
    list_display = ("username","phoneno","emailid" ,"datetime")

admin.site.register(Userdata,UserdataAdmin)